import Button from "../../ui/Button.tsx";
import Popup from "../../ui/Popup.tsx";

interface SuccessPopupProps {
    isOpen: boolean;
    onContinue: () => void;
}

const AuthSuccessPopup = ({ isOpen, onContinue }: SuccessPopupProps) => (
    <Popup isOpen={isOpen} onClose={onContinue}>
        <p>You've successfully connected your osu! account.</p>
        <p className="mt-2">Wait for 2-5 minutes and navigate to users page.</p>
        <Button
            onClick={onContinue}
            className="text-xl rounded-md p-4 bg-green-800 mt-4 w-1/4 hover:bg-green-900"
            content="OK"
        />
    </Popup>
);

export default AuthSuccessPopup;