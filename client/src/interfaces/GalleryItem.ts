export default interface GalleryItemProps {
  img: HTMLImageElement;
  likes: number;
  likeHandler: (...args: string[])=>any
}
