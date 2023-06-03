import { GalleryItem } from "./GalleryItem";
import { MasonryGrid } from "./MasonryGrid";

export function Images({ images }: { images: HTMLImageElement[] }) {
  return (
    <MasonryGrid>
      {images.map((img, id) => (
        <GalleryItem img={img} key={id} />
      ))}
    </MasonryGrid>
  );
}
