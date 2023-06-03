import { useState } from "react";
import { NavItem as Item } from "../interfaces/navbar";
import NavItem from "./NavItem";
import Logo from "./Logo";
import { AiOutlineClose, AiOutlineMenu } from "react-icons/ai";

export default function Navbar() {
  const [isOpen, setOpen] = useState<boolean>(false);

  const navbarToggler = () => {
    setOpen((isOpen) => !isOpen);
  };
  const navItems: Item[] = [
    {
      to: "/",
      text: "Home",
    },
    {
      to: "/gallery",
      text: "Gallery",
    },
    {
      to: "/support",
      text: "Support Us",
    },
  ];

  return (
    <nav className="min-h-[8vh] px-8 pt-4 flex lg:flex-row lg:mr-auto  items-center">
      <Logo />
      <div className={"rwd-navbar " + (isOpen ? "active" : "")}>
        {navItems.length &&
          navItems.map((navItem, id) => (
            <NavItem {...navItem} onclick={(isOpen) => setOpen(!isOpen)} key={id}/>
          ))}
      </div>
      <span className="lg:hidden ml-auto" onClick={navbarToggler}>
        {isOpen ? <AiOutlineClose /> : <AiOutlineMenu />}
      </span>
    </nav>
  );
}
