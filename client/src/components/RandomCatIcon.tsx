import blackcat from "../assets/blackcat-icon.png";
import whiteCatSm from "../assets/sm-whitecat-icon.png";
import whiteCat from "../assets/whitecat-icon.png";
import grayCat from "../assets/graycat-icon.png";

export function RandomCatIcon() {
  let cats = [blackcat, whiteCat, whiteCatSm, grayCat];

  return <img src={cats[Math.floor(Math.random() * cats.length)]} />;
}
