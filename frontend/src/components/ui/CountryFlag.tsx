import {getCountryName, isValidCountryCode} from "../../utils/country.ts";

interface CountryFlagProps {
    countryCode: string;
    className?: string;
    showName?: boolean;
    nameClassName?: string;
}

const CountryFlag = ({countryCode, className, showName, nameClassName}: CountryFlagProps) => {
    if (!isValidCountryCode(countryCode)) {
        return null;
    }

    const code = countryCode.toLowerCase();
    const countryName = getCountryName(countryCode);

    return (
        <span className={`inline-flex items-center gap-2 h-3.5 md:h-4 ${className ?? ""}`}>
            <img
                src={`/flags/4x3/${code}.svg`}
                alt={countryName ?? countryCode.toUpperCase()}
                title={countryName ?? countryCode.toUpperCase()}
                className="w-5 md:w-6 h-full shrink-0 rounded-sm object-cover"
                loading="lazy"
                decoding="async"
            />
            {showName && countryName && (
                <span className={`leading-none ${nameClassName ?? "text-xs text-zinc-400"}`}>
                    {countryName}
                </span>
            )}
        </span>
    );
};

export default CountryFlag;
