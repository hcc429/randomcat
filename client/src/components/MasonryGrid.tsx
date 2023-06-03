import { ReactNode } from "react";
import Masonry, { ResponsiveMasonry } from "react-responsive-masonry";

export function MasonryGrid({ children }: { children: ReactNode }) {
  return (
    <ResponsiveMasonry
      columnsCountBreakPoints={{ 350: 1, 768: 2, 996: 3 }}
      className="mt-12 p-3 mx-3"
    >
      <Masonry gutter="1rem">{children}</Masonry>
    </ResponsiveMasonry>
  );
}
