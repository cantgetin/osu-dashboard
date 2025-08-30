interface ButtonProps {
    keyNumber?: number;
    onClick: (key: number) => void;
    className: string;
    content: React.ReactNode;
}

const MyButton = (props: ButtonProps) => {
    return (
        <button
            key={props.keyNumber}
            onClick={() => props.onClick(props.keyNumber != undefined ? props.keyNumber : 1)}
            className={props.className}
        >
            {props.content}
        </button>
    );
};

export default MyButton;