import React from 'react';

interface ListProps<T> {
    items: T[];
    renderItem: (item: T) => React.ReactNode
    className?: string
    title?: React.ReactNode
}

export default function List<T>(props: ListProps<T>) {
    return (
        <>
            {props.items?.length > 0 && (
                <div className={props.className}>
                    {props.title && (
                        <div className="h-1 text-sm text-gray-500">
                            {props.title}
                        </div>
                    )}
                    {props.items.map(props.renderItem)}
                </div>
            )}
        </>
    );
}