import useLikeBuffer from "../hooks/useLikes";
import GalleryItemProps from "../interfaces/GalleryItem";
import { GalleryItem } from "./GalleryItem";
import { MasonryGrid } from "./MasonryGrid";

export function Images({ images }: { images: GalleryItemProps[] }) {
  const insertLike = useLikeBuffer();
  return (
    <MasonryGrid>
      {images.map((img, id) => (
        <GalleryItem
          {...img}
          key={id}
          likeHandler={() => insertLike(img.img.src)}
        />
      ))}
    </MasonryGrid>
  );
}
