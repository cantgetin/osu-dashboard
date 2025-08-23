interface ContainerProps {
    children: React.ReactNode
}

const Container = (props: ContainerProps) => {
    return (
        <div className="flex flex-col gap-3 md:gap-5 bg-zinc-900 p-2 md:p-4 rounded-lg w-full">
            {props.children}
        </div>
    );
};

export default Container;