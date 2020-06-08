import React, {
  FunctionComponent,
  useEffect,
  useRef,
  useState,
  useCallback,
  useMemo,
  RefObject,
} from 'react';

import { CardProps, DivRef } from './general';
import { EndpointCardLayer } from './EndpointCardBase';
import { EndpointCardLabels } from './EndpointCardLabels';
import { EndpointCardHeader } from '~/components/EndpointCardHeader';
import { AccessPoint, CenterGetter } from '~/components/AccessPoint';

import { Vec2 } from '~/domain/geometry';
import { ServiceCard } from '~/domain/service-card';
import { Link, AccessPoint as AccessPointDatum } from '~/domain/service-map';
import { ids } from '~/domain/ids';
import { Dictionary } from '~/domain/misc';
import { sizes } from '~/ui';

import css from './styles.scss';

export type Props = CardProps & {
  onHeaderClick?: (card: ServiceCard) => void;
  onEmitAPConnectorCoords?: (apId: string, coords: Vec2) => void;
};

export const Component: FunctionComponent<Props> = props => {
  const [divRef, setContentRef] = useState<DivRef>(null as any);
  const centerGetters = useMemo((): Map<string, CenterGetter> => {
    return new Map();
  }, []);

  // prettier-ignore
  const onEmitContentRef = useCallback((ref: DivRef) => {
    setContentRef(ref);
  }, [setContentRef]);

  const onConnectorReady = useCallback((apId: string, cg: CenterGetter) => {
    centerGetters.set(apId, cg);
  }, []);

  const emitConnectorCoords = useCallback(() => {
    // prettier-ignore
    if (props.onEmitAPConnectorCoords == null || divRef == null) return;

    const bbox = divRef.current!.getBoundingClientRect();
    if (bbox.width === 0 || bbox.height === 0) {
      return;
    }

    // console.log(`emitting connector coords from card: ${props.card.appLabel || props.card.id}`);
    // console.log(`center getters: `, centerGetters);
    centerGetters.forEach((cg: CenterGetter, apId: string) => {
      const connectorCenter = centerGetters.get(apId)!();
      const relCoords = connectorCenter.sub(Vec2.fromXY(bbox));

      const factorCoords = Vec2.from(
        relCoords.x / bbox.width,
        relCoords.y / bbox.height,
      );

      const svgCoords = Vec2.from(
        factorCoords.x * props.coords.w,
        factorCoords.y * props.coords.h,
      ).add(Vec2.fromXY(props.coords));

      props.onEmitAPConnectorCoords!(apId, svgCoords);
    });
  }, [props.onEmitAPConnectorCoords, divRef, centerGetters, props.coords]);

  // react to placement change
  useEffect(emitConnectorCoords, [
    props.coords,
    centerGetters,
    emitConnectorCoords,
  ]);

  const accessPoints = useMemo(() => {
    if (!props.accessPoints) return [];
    const aps = Array.from(props.accessPoints!.values());

    return aps.map((ap: AccessPointDatum) => {
      return (
        <AccessPoint
          key={ap.id}
          id={ap.id}
          port={ap.port}
          protocol={ap.protocol}
          onConnectorReady={onConnectorReady}
        />
      );
    });
  }, [props.accessPoints, onConnectorReady]);

  // prettier-ignore
  const onHeightChange = useCallback((card: ServiceCard, h: number) => {
    if (props.onHeightChange) props.onHeightChange(card, h);

    // WARN: do not emit new connector coords from here cz
    // old coords will be received in props.coords hence wrong connector coords
    // will be emitted

  }, [emitConnectorCoords, props.onHeightChange]);

  return (
    <EndpointCardLayer
      {...props}
      onHeightChange={onHeightChange}
      onEmitContentRef={onEmitContentRef}
    >
      <EndpointCardHeader card={props.card} onClick={props.onHeaderClick} />

      {accessPoints.length > 0 && (
        <div className={css.accessPoints}>{accessPoints}</div>
      )}

      {props.active && <EndpointCardLabels labels={props.card.labels} />}
    </EndpointCardLayer>
  );
};

Component.displayName = 'EndpointCardContent';
export const EndpointCardContent = React.memo(Component);