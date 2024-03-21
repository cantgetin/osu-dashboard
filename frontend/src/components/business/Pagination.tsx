import List from "../logic/List.tsx";
import Button from "../ui/Button.tsx";

interface PaginationProps {
    className?: string;
    onPageChange: (page: number) => void;
    pages: number;
    currentPage: number
}

const Pagination = (props: PaginationProps) => {
    const buttons = Array.apply(null, Array(props.pages)).map(function (_, i) {
        return i + 1
    })

    return (
            <List className={props.className}
                  title="Page:"
                  items={buttons}
                  renderItem={(num) =>
                      <Button keyNumber={num}
                              key={num}
                              onClick={(key: number) => {
                                  props.onPageChange(key)
                              }}
                              className={"rounded-md w-12 " + (num === props.currentPage ? "bg-white text-black" : "bg-zinc-800")}
                              content={num.toString()}
                      />
                  }
            />
    );
};

export default Pagination;