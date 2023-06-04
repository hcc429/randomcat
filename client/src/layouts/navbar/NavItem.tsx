import { NavLink } from "react-router-dom";
import NavItemProps from "../../interfaces/Navbar";

export default function NavItem({ to, onclick, text }: NavItemProps) {
  return (
    <div className="nav-item" onClick={onclick ?? undefined}>
      <NavLink
        to={to}
        className={({ isActive }) =>
          (isActive ? "bg-amber-100" : "hover:bg-amber-50") +
          " text-2xl transition duration-300 px-4 py-3 rounded-md hover:shadow-sm shadow-amber-200"
        }
      >
        {text}
      </NavLink>
    </div>
  );
}
