export default function Footer() {
  return (
    <footer className="bg-primary text-slate-100  p-5 min-h-[5rem] text-center tracking-wider">
      <h2>
        Assets from{" "}
        <a href="https://unsplash.com" className="link">
          Unsplash
        </a>
        ,
        <a href="https://www.pexels.com/" className="link">
          pexel
        </a>
        ,
        <a href="https://icons8.com" className="link">
          icon8
        </a>
        <br />
        Welcome to share your cat with us!
      </h2>
      <h2 className="text-white text-lg">
        {new Date().getFullYear()} randomcat &copy;
      </h2>
    </footer>
  );
}
