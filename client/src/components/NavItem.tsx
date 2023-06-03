import { NavLink } from "react-router-dom";
import { NavItem } from "../interfaces/navbar";

export default function NavItem(props: NavItem) {
  return (
    <div className="nav-item" onClick={props.onclick ?? undefined}>
      <NavLink
        to={props.to}
        className={({ isActive }) =>
          (isActive ? "bg-amber-100" : "hover:bg-amber-50") +
          " text-2xl transition duration-300 px-4 py-3 rounded-md hover:shadow-sm shadow-amber-200"
        }
      >
        {props.text}
      </NavLink>
    </div>
  );
}
