import {ReactNode} from 'react';

interface PopupProps {
    isOpen: boolean;
    onClose: () => void;
    children: ReactNode;
    title?: string;
}

// TODO: look into that, idfk how it works
const Popup = ({isOpen, onClose, children, title = "Authorization Successful"}: PopupProps) => {
    if (!isOpen) return null;

    return (
        <div
            className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50"
            onClick={(e) => e.target === e.currentTarget && onClose()}
        >
            <div className="bg-zinc-900 rounded-lg w-full max-w-[600px] min-w-[200px] max-h-[50vh] overflow-y-auto">
                <div className="flex justify-between items-center p-6">
                    <div className="flex items-center">
                        <svg className="w-8 h-8 text-green-500 mr-3" fill="none" stroke="currentColor"
                             viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={3} d="M5 13l4 4L19 7"/>
                        </svg>
                        <h2 className="text-xl font-medium">{title}</h2>
                    </div>
                    <button
                        onClick={onClose}
                        className="text-gray-500 hover:text-gray-700 text-4xl px-4 py-2"
                    >
                        &times;
                    </button>
                </div>

                <div className="p-6 flex flex-col items-center justify-center text-center">
                    {children}
                </div>
            </div>
        </div>
    );
};

export default Popup;