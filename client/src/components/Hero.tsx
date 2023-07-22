import herobg from "../assets/hero-bg.png";

function LeftComponent() {
  return (
    <div className="bg-primary w-full   p-6 rounded-md font-handwriting text-white md:max-w-[20rem] xl:max-w-[28rem] z-10 md:mr-[-3%]">
      {/* <img src={logo} alt="" className="w-16" /> */}

      <div className="flex gap-4 justify-center">
        <h2 className="text-3xl overflow-hidden lg:text-4xl">
          <span className="block hero-animate-left-text ">Develop</span>
        </h2>
        <h2 className="text-3xl overflow-hidden lg:text-4xl">
          <span className="block hero-animate-center-text ">With</span>
        </h2>
        <h2 className="text-3xl overflow-hidden lg:text-4xl">
          <span className="block hero-animate-right-text ">Cat</span>
        </h2>
      </div>
    </div>
  );
}

function RightComponent() {
  return (
    <div>
      <img src={herobg} alt="" className="w-full max-w-[72rem] hero-bg " />
    </div>
  );
}

export default function Section1() {
  return (
    <div className="p-10 flex flex-col min-h-[20rem] md:min-h-[40rem] items-center md:flex-row md:justify-end md:pr-0">
      <LeftComponent />
      <RightComponent />
    </div>
  );
}
