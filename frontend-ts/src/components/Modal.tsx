import { Dialog } from "@headlessui/react";
import React from "react";

export function Modal({
  title,
  children,
  close,
}: {
  title: string;
  children: React.ReactElement;
  close: () => void;
}) {
  return (
    <Dialog
      as="div"
      className="fixed inset-0 z-10 overflow-y-auto"
      open
      onClose={() => close()}
    >
      <div className="min-h-screen px-4 text-center">
        <Dialog.Overlay className="fixed inset-0 bg-gray-500 bg-opacity-40" />
        <div className="inline-block w-full max-w-md mt-36 px-5 pt-4 pb-5 overflow-hidden text-left transform bg-white rounded">
          <Dialog.Title as="p" className="text-lg font-semibold mb-2">
            {title}
          </Dialog.Title>
          {children}
        </div>
      </div>
    </Dialog>
  );
}
