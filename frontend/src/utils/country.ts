export function isValidCountryCode(countryCode: string): boolean {
    if (!countryCode || countryCode.length !== 2 || countryCode === "XX") {
        return false;
    }

    return /^[A-Za-z]{2}$/.test(countryCode);
}

const regionNames = new Intl.DisplayNames(["en"], {type: "region"});

export function getCountryName(countryCode: string): string | null {
    if (!isValidCountryCode(countryCode)) {
        return null;
    }

    return regionNames.of(countryCode.toUpperCase()) ?? countryCode.toUpperCase();
}
