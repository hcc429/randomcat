import { useState } from "react";


export default function GuideSection() {
  const [width, setWidth] = useState(500);
  const [height, setHeight] = useState(400);

  const widthHandler = (e: any) => {
    setWidth((_) => parseInt(e.target.value) || 500);
  };
  const heightHandler = (e: any) => {
    setHeight((_) => parseInt(e.target.value) || 400);
  };

  const maxLength = 5000;
  const isValid = (len: number) => {
    return len > 0 && len <= maxLength;
  };
  const getUrl = () => {
    return `https://randomcat.io/api/image?w=${width}&h=${height}`;
  };
  return (
    <div className="mt-6 lg:mt-8 py-8 lg:py-20 bg-secondary text-white">
      <h2 className="text-3xl lg:text-4xl font-handwriting text-center ">
        Feel free to use!
      </h2>

      <div className="text-center mt-4 py-4 ">
        {isValid(width) && isValid(height) ? (
          <>
            
            <code className="text-xs md:text-base py-4 px-1 md:px-3 lg:px-6 bg-gray-50 rounded-sm text-blue-500">
              <a href={getUrl()}>{getUrl()}</a>
            </code>
          </>
        ) : (
            <h3 className="text-sm lg:text-xl font-bold text-red-500">
                /ᐠ｡ꞈ｡ᐟ\ Meow 0 Meow Meow {maxLength}...
            </h3>
        )}
      </div>
      <div className="text-md lg:text-2xl flex flex-col  gap-3 items-center m-5 p-5">
        <input
          type="number"
          className={"input-number" + (isValid(width) ? "" : " invalid-input")}
          onChange={widthHandler}
          placeholder="width"
        />
        <input
          type="number"
          className={"input-number" + (isValid(height) ? "" : " invalid-input")}
          onChange={heightHandler}
          placeholder="height"
        />
      </div>
    </div>
  );
}
