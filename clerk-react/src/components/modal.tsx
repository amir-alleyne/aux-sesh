import React from 'react'
import {Modal, Button, Spin} from "antd";

type GenericModalProps = {
    title: React.ReactNode;
    open: boolean;
    loading?: boolean;
    onOk?: () => void;
    onClose: () => void;         // Callback to handle modal closing.
    onReload?: () => void;       // Optional callback for a "reload" action.
    children: React.ReactNode;
  };


export default function GenModal({
    title,
    open,
    loading,
    onOk,
    onClose,
    onReload,
    children,
}: GenericModalProps) {
  return (
    <Modal
      title={title}
      open={open}
      onOk={onOk ? onOk : onClose}
      onCancel={onClose}  
      footer={
        onReload && (
          <Button type="primary" onClick={onReload}>
            Reload
          </Button>
        )
      }
    >
      {loading ? <Spin spinning={loading}>{children}</Spin> : children}
    </Modal>
  )
}
