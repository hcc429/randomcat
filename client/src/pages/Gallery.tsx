import { Images } from "../components/Images";
import { useEffect} from "react";
import useImageLoader from "../hooks/useImageLoader";
import { ImageSkeleton } from "../components/Skeleton";
import { MasonryGrid } from "../components/MasonryGrid";

export default function Gallery() {
  const limit = 6;
  let { images, isLoading, nextPage, isEnd } = useImageLoader(limit);
  useEffect(() => {
    nextPage();
  }, []);
  return (
    <div className="m-16  mx-16 max-w-[80rem] lg:mx-auto">
      <Images images={images} />
      {isLoading && (
        <MasonryGrid >
          {[...Array(limit)].map((i, id) => (
            <ImageSkeleton key={id} />
          ))}
        </MasonryGrid>
      )}
      <div className="text-center m-9">
        {isEnd ? (
          <h2 className="text-2xl "> No more :(</h2>
        ) : (
          <button className="btn btn-primary" onClick={nextPage}>
            Load More!
          </button>
        )}
      </div>
    </div>
  );
}
