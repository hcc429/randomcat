import { useState } from "react";
import ImageService from "../services/image.service";
import GalleryItemProps from "../interfaces/GalleryItem";
const loadImage = (src: string, likes: number) => {
  return new Promise((resolve, reject) => {
    const img = new window.Image();
    img.src = src;
    img.onload = () => resolve({img, likes});
    img.onerror = reject;
  });
};

const useImageLoader = (limit: number) => {
  const [images, setImages] = useState<GalleryItemProps[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [isLoading, setIsLoading] = useState(false);
  const [isEnd, setEnd] = useState(false);
  const nextPage = async () => {
    setIsLoading(true);
    try {
      let nextImages = await ImageService.getPage(currentPage, limit);
      if (nextImages && nextImages.length) {
        let newImages = (await Promise.all(
          nextImages.map((img) => loadImage(img.URL, img.Likes))
        )) as GalleryItemProps[];
        setImages((i) => [...i, ...newImages]);
        setCurrentPage((prevPage) => prevPage + 1);
      } else if (nextImages.length == 0) {
        setEnd(true);
      }
    } catch (error) {
      console.log(error);
    } finally {
      setIsLoading(false);
    }
  };

  return { images, nextPage, isLoading, isEnd };
};

export default useImageLoader;
