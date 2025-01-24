import { useMemoizedFn } from "ahooks";
import { Modal, ModalProps } from "antd";
import { useMemo, useState } from "react";

export type ModalOperation = {
    open: () => void;
    close: () => void;
    isOpen: boolean;
}

export type ModalChildrenProps<T = {}> = (ModalOperation & T);

export function useModal<T>(childrenFn: (props: ModalChildrenProps<T>) => React.ReactNode, childrenProps: T) {
    const [isOpen, setIsOpen] = useState<boolean>(false);
    const open = useMemoizedFn(() => setIsOpen(true));
    const close = useMemoizedFn(() => setIsOpen(false));
    const gatheredProps = useMemo(() => ({ open, close, isOpen, ...childrenProps }), [isOpen, childrenProps]);
    const modal = (props: ModalProps) => <Modal {...props} open={isOpen}>
        {childrenFn(gatheredProps)}
    </Modal>;
    return { modal, open, close, isOpen }
}