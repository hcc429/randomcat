interface GalleryItemProps {
  img: HTMLImageElement;
}

export function GalleryItem({ img }: GalleryItemProps) {
  return (
    <div className="overflow-hidden rounded-md gallery-item">
      <img
        src={img.src}
        className="w-full brightness-90 hover:brightness-100 hover:scale-105 transition-all duration-300"
      />
    </div>
  );
}
