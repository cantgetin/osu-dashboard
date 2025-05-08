import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import Layout from "../components/ui/Layout.tsx";
import AuthSuccessPopup from "../components/features/common/AuthSuccessPopup.tsx";
import useOsuAuth from "../hooks/useOsuAuth.ts";

const Authorize = () => {
    const { isSuccess, handlePopupContinue } = useOsuAuth();

    return (
        <Layout title="Authorize">
            <LoadingSpinner />
            <AuthSuccessPopup
                isOpen={isSuccess}
                onContinue={handlePopupContinue}
            />
        </Layout>
    );
};

export default Authorize;