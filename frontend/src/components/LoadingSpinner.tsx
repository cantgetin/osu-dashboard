import React from 'react';
import { SpinnerCircularFixed } from 'spinners-react';

const LoadingSpinner = () => {
    return (
        <div className="flex justify-center items-center h-screen w-screen">
            <SpinnerCircularFixed size={90} thickness={127} speed={112} color="#2e2e2e"/>
        </div>
    );
}

export default LoadingSpinner;