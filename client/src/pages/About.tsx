import aboutbg from "../assets/about-bg.png";
import { RandomCatIcon } from "../components/RandomCatIcon";
import { AiFillGithub, AiFillMail } from "react-icons/ai";
import { ReactElement } from "react";

interface IAboutOption {
  icon: ReactElement;
  text: string;
  link: string;
}
const AboutOption = ({ icon, text, link }: IAboutOption) => {
  return (
    <div className="min-w-[18rem]    mx-auto p-6 rounded-xl border-dotted border-4 border-primary">
      <a
        href={link}
        className="text-3xl  lg:text-4xl flex items-center gap-3  font-handwriting "
      >
        {icon}
        {text}
      </a>
    </div>
  );
};
export default function About() {
  const aboutOptions: IAboutOption[] = [
    {
      icon: <AiFillMail></AiFillMail>,
      text: "Contact Us",
      link: "mailto:israndom.cat@gmail.com",
    },
    {
      icon: <AiFillGithub></AiFillGithub>,
      text: "Give us a ☆!",
      link: "https://github.com/hcc429/randomcat",
    },
  ];
  return (
    <div className="mt-4 p-4 lg:mt-8 lg:p-6  ">
      <div>
        <img
          src={aboutbg}
          className="min-w-[90%] xl:max-w-[64rem] mx-auto"
          alt=""
        />
      </div>
      <div className="mt-16 p-3 mb-40 md:m-6 md:p-6 flex flex-col gap-12 items-center">
        {aboutOptions && aboutOptions.map((opt) => <AboutOption {...opt} />)}
      </div>
      <div className="w-16 mb-[-1rem] lg:mb-[-1.5rem]">
        <RandomCatIcon />
      </div>
    </div>
  );
}
