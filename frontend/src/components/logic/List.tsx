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
            {props.items != null || undefined ?
                <>
                    {props.items.length > 0 && props.title}
                    {props.items.length > 0 ? props.items.map(props.renderItem) : null}
                </>
                : null}

        </div>
    )
}