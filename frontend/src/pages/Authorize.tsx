import {useEffect} from "react";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import {redirectToAuthorize} from "../utils/utils.ts";
import {TokenResponse} from "../interfaces/TokenResponse.ts";
import axios from "axios";
import * as queryString from "querystring";
import {useNavigate} from "react-router-dom";

async function handleOsuSiteRedirect(state: string, code: string) {
    console.log(`redirect state: ${state} local state: ${localStorage.getItem('state')}, all good`)
    if (state == localStorage.getItem('state')) {
        localStorage.setItem('code', code?.toString())
        console.log('set the code to local storage, now exchange code for token')

        // exchange code for authorization token
        let res = await axios.post('/api/exchange', { code: code })

        let data: TokenResponse = res.data
        localStorage.setItem('access_token', data.access_token)
        localStorage.setItem('refresh_token', data.refresh_token)
        console.log('set access and refresh token to localstorage', res)
    }
}

const Authorize = () => {
    const navigate = useNavigate();

    const authorize = async () => {
        if (localStorage.getItem('access_token') != null) {
            navigate("/maps");
        } else {
            redirectToAuthorize();
        }
    };

    useEffect(() => {
        const { search } = window.location;
        const { code, state } = queryString.parse(search);

        if (code?.toString() != undefined && state?.toString() != undefined) {
            handleOsuSiteRedirect(state.toString(), code.toString()).then(() => {
                navigate("/users");
            })
        }
        else {
            authorize().then(r => console.log(r));
        }
    }, []);

    return (
        <LoadingSpinner/>
    );
};

export default Authorize;