import GalleryItemProps from "../interfaces/GalleryItem";
import { useState } from "react";
import { AiFillHeart } from "react-icons/ai";

export function GalleryItem({
  img,
  likes: likeNums,
  likeHandler: parentLikeHandler,
}: GalleryItemProps) {
  let [likes, setLikes] = useState(likeNums);

  const likeHandler = () => {
    parentLikeHandler();
    setLikes((likes) => likes + 1);
  };

  return (
    <div className="gallery-item select-none border-2 rounded-md ">
      <div className="overflow-hidden">
        <img
          src={img.src}
          className="w-full brightness-95 hover:brightness-100 hover:scale-[1.02] transition-all duration-300"
        />
      </div>
      <div className=" rounded-md rounded-tr-none rounded-tl-none p-3 flex gap-5 items-center">
        <div
          className="cursor-pointer active:scale-75 transition-all duration-75"
          onClick={likeHandler}
        >
          <AiFillHeart className="text-red-400" />
        </div>
        {likes}
      </div>
    </div>
  );
}
