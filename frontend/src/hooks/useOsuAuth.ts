import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import queryString from "query-string";
import { handleOsuSiteRedirect, redirectToAuthorize } from "../utils/utils";

const useOsuAuth = () => {
    const navigate = useNavigate();
    const [isSuccess, setIsSuccess] = useState(false);

    const authorize = async () => {
        redirectToAuthorize();
    };

    const handlePopupContinue = () => {
        setIsSuccess(false);
        navigate("/users");
    };

    useEffect(() => {
        const { search } = window.location;
        const { code, state } = queryString.parse(search);

        if (code?.toString() && state?.toString()) {
            handleOsuSiteRedirect(state.toString(), code.toString())
                .then(() => setIsSuccess(true))
                .catch(error => {
                    console.error("Authorization failed:", error);
                    navigate("/");
                });
        } else {
            authorize();
        }
    }, [navigate]);

    return { isSuccess, handlePopupContinue };
};

export default useOsuAuth;