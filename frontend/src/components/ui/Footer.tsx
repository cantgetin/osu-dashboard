import FooterLink from "./FooterLink.tsx";

const Footer = () => {
    return (
        <div className="z-20 bg-zinc-900 w-full bg-opacity-85 text-[16px] sm:text-[16px]
        backdrop-blur-sm flex justify-center items-center">
            <div className="max-w-[1152px] w-full px-4 h-14 flex items-center justify-center flex-wrap">
                <FooterLink href="https://osu.ppy.sh/users/7192129">Creator</FooterLink>
                <FooterLink href="https://github.com/cantgetin/osu-dashboard">Source code</FooterLink>
                <FooterLink href="/feedback" target="_self">Feedback</FooterLink>
            </div>
        </div>
    );
};

export default Footer;