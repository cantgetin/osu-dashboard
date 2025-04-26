import {useEffect, useState} from "react";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import {redirectToAuthorize} from "../utils/utils.ts";
import axios from "axios";
import queryString from "query-string";
import {useNavigate} from "react-router-dom";
import Popup from "../components/ui/Popup.tsx";
import Button from "../components/ui/Button.tsx";

async function handleOsuSiteRedirect(state: string, code: string) {
    console.log(`redirect state: ${state} local state: ${localStorage.getItem('state')}, all good`)
    if (state == localStorage.getItem('state')) {
        localStorage.setItem('code', code?.toString())
        console.log('set the code to local storage, now exchange code for token')

        // create following with obtained code
        await axios.post(`/api/following/create/${code}`);
    }
}

const Authorize = () => {
    const navigate = useNavigate();
    const [showSuccessPopup, setShowSuccessPopup] = useState(false);

    const authorize = async () => {
        redirectToAuthorize();
    };

    useEffect(() => {
        const { search } = window.location;
        const { code, state } = queryString.parse(search);

        if (code?.toString() != undefined && state?.toString() != undefined) {
            handleOsuSiteRedirect(state.toString(), code.toString()).then(() => {
                setShowSuccessPopup(true); // Show popup instead of immediate navigation
            }).catch(error => {
                console.error("Authorization failed:", error);
                navigate("/"); // Redirect to home if there's an error
            });
        }
        else {
            authorize().then(r => console.log(r));
        }
    }, []);

    const handlePopupContinue = () => {
        setShowSuccessPopup(false);
        navigate("/users");
    };

    return (
        <>
            <LoadingSpinner/>
            <Popup isOpen={showSuccessPopup} onClose={() => handlePopupContinue()}>
                <p>You've successfully connected your osu! account.</p>
                <p className="mt-2">Wait for 2-5 minutes and navigate to users page.</p>
                <Button onClick={() => handlePopupContinue()}
                        className="text-xl rounded-md p-4 bg-green-800 mt-4 w-1/4 hover:bg-green-900"
                        content="OK"
                />
            </Popup>
        </>
    );
};

export default Authorize;