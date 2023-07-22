export default interface NavItemProps {
  to: string;
  text: string;
  onclick?: (...args: any[]) => any;
}
