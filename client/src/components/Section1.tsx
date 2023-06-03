import bg1 from "../assets/bg1.png";
import { Timeline, Tween } from "react-gsap";

function LeftComponent() {
  return (
    <div className="bg-primary w-full   p-6 rounded-md font-handwriting text-white md:max-w-[20rem] xl:max-w-[28rem] z-10 md:mr-[-3%]">
      {/* <img src={logo} alt="" className="w-16" /> */}

      <div className="flex gap-4 justify-center">
        <h2 className="text-3xl overflow-hidden lg:text-4xl">
          <Tween
            from={{ x: "100%", opacity: 0.5 }}
            to={{ x: 0, opacity: 1 }}
            ease="Power4.easeOut"
            position="<30%"
          >
            <span className="block">Develop</span>
          </Tween>
        </h2>
        <h2 className="text-3xl overflow-hidden lg:text-4xl">
          <Tween
            from={{ y: "100%", opacity: 0.5 }}
            to={{ y: 0, opacity: 1 }}
            ease="Power4.easeOut"
            position="<50%"
          >
            <span className="block">With</span>
          </Tween>
        </h2>
        <h2 className="text-3xl overflow-hidden lg:text-4xl">
          <Tween
            from={{ x: "-100%", opacity: 0.5 }}
            to={{ x: 0, opacity: 1 }}
            position={"<-25%"}
            ease="Power4.easeOut"
          >
            <span className="block">Cat</span>
          </Tween>
        </h2>
      </div>
    </div>
  );
}

function RightComponent() {
  return (
    <div>
      <Tween
        from={{ scale: 1.4 }}
        to={{ scale: 1 }}
        position={0}
        ease="elastic.out(1.5,1)"
        duration={2.5}
      >
        <img src={bg1} alt="" className="w-full max-w-[72rem] " />
      </Tween>
    </div>
  );
}

export default function Section1() {
  return (
    <div className="p-10 flex flex-col  min-h-[40rem] items-center md:flex-row md:justify-end md:pr-0">
      <Timeline>
        <LeftComponent />
        <RightComponent />
      </Timeline>
    </div>
  );
}
