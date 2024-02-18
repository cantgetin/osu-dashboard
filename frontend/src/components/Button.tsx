interface ButtonProps {
    keyNumber: number;
    onClick: (key: number) => void;
    className: string;
    content: string;
}

const Button = (props: ButtonProps) => {
    return (
        <button
            key={props.keyNumber}
            onClick={() => props.onClick(props.keyNumber)}
            className={props.className}
        >
            {props.content}
        </button>
    );
};

export default Button;