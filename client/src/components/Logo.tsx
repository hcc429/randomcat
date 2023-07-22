import logo from "../assets/logo-brown.png";

export default function Logo() {
  return (
    <div className="flex items-center gap-4">
      <img src={logo} alt="" className="w-16" />
      <h1 className="text-3xl md:text-4x lg:text-5xl text-primary font-handwriting ">Random Cat</h1>
    </div>
  );
}
