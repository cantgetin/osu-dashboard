import aveta from "aveta";

interface StatsDifferenceProps {
    difference: number;
    className?: string;
    forceShowDiff?: boolean;
}

const StatsDifference = (props: StatsDifferenceProps) => {
    return (
        <>
            {props.difference != 0 || props.forceShowDiff ?
                <div className={`flex gap-2 items-center w-full justify-end ${props.className}`}>
                    {props.difference >= 0 ?
                        <h1 className="text-xs">▲</h1>
                        :
                        <h1 className="text-xs">▼</h1>
                    }
                    <h1 className="text-2xl">{aveta(props.difference)}</h1>
                </div>
                : null}
        </>

    );
};

export default StatsDifference;