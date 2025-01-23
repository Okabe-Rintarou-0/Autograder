import { useMemoizedFn } from "ahooks";
import { Modal, ModalProps } from "antd";
import { useMemo, useState } from "react";

export type ModalChildrenProps<T> = {
    open: () => void;
    close: () => void;
    isOpen: boolean;
    childrenProps?: T
}

export function useModal<T>(childrenFn: (props: ModalChildrenProps<T>) => React.ReactNode) {
    const [isOpen, setIsOpen] = useState<boolean>(false);

    const open = useMemoizedFn(() => setIsOpen(true));
    const close = useMemoizedFn(() => setIsOpen(false));
    const childrenProps = useMemo(() => ({ open, close, isOpen }), [isOpen]);
    const modal = (props: ModalProps) => <Modal {...props} open={isOpen}>
        {childrenFn(childrenProps)}
    </Modal>;
    return { modal, open, close, isOpen }
}