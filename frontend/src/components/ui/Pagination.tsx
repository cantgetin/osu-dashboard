import List from "../logic/List.tsx";
import MyButton from "./MyButton.tsx";

interface PaginationProps {
    className?: string;
    onPageChange: (page: number) => void;
    pages: number;
    currentPage: number
}

const Pagination = (props: PaginationProps) => {
    if (props.pages <= 1) {
        return null;
    }

    let startPage: number, endPage: number;
    const maxPagesToShow = 10;

    if (props.pages <= maxPagesToShow) {
        startPage = 1;
        endPage = props.pages;
    } else {
        if (props.currentPage <= 5) {
            startPage = 1;
            endPage = maxPagesToShow;
        } else if (props.currentPage + 4 >= props.pages) {
            startPage = props.pages - maxPagesToShow + 1;
            endPage = props.pages;
        } else {
            startPage = props.currentPage - 5;
            endPage = props.currentPage + 4;
        }
    }

    const buttons = [];
    for (let i = startPage; i <= endPage; i++) {
        buttons.push(i);
    }

    return (
        <List className={props.className}
              title="Page:"
              items={buttons}
              renderItem={(num) =>
                  <MyButton keyNumber={num}
                            key={num}
                            onClick={(key: number) => {
                              props.onPageChange(key)
                          }}
                            className={"rounded-md w-12 " + (num === props.currentPage ? "bg-white text-black"
                              : "bg-zinc-800")}
                            content={num.toString()}
                  />
              }
        />
    );
};

export default Pagination;