import React from 'react';

interface ListProps<T> {
    items: T[];
    renderItem: (item: T) => React.ReactNode
    className?: string
    title?: React.ReactNode
}

export default function List<T>(props: ListProps<T>) {
    return (
        <div className={props.className}>
            {props.items?.length > 0 && (
                <div className="space-y-1">
                    {props.title && (
                        <div className="text-sm text-gray-500 px-1">
                            {props.title}
                        </div>
                    )}
                    {props.items.map(props.renderItem)}
                </div>
            )}
        </div>
    );
}