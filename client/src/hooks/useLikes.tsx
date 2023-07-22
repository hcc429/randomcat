import { useState, useEffect, useRef } from "react";
import Like from "../interfaces/Like";
import ImageService from "../services/image.service";

const useLikeBuffer = (interval = 3000) => {
  let [likeBuffer, setLikeBuffer] = useState<Like[]>([]);
  let bufferRef = useRef(likeBuffer);
  bufferRef.current = likeBuffer;

  let combineLikes = () => {
    return bufferRef.current.reduce((result, cur) => {
      const existingItem = result.find((item) => item.url == cur.url);
      if (existingItem) {
        existingItem.amount += cur.amount;
      } else result.push(cur);
      return result;
    }, [] as Like[]);
  };

  let updateLike = () => {
    if (bufferRef.current.length == 0) return;
    let combinedLikes = combineLikes();
    ImageService.likeImages(combinedLikes);
    setLikeBuffer([]);
  };

  let insertLike = (likesUrl: string) => {
    setLikeBuffer((likeBuffer) => [
      ...likeBuffer,
      { url: likesUrl, amount: 1 },
    ]);
  };
  useEffect(() => {
    let intervalId = setInterval(updateLike, interval);

    return () => {
      clearInterval(intervalId);
    };
  }, []);

  return insertLike;
};

export default useLikeBuffer;
