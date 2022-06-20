import { Button } from "@mantine/core";
import Link from "next/link";

export type WaitListButtonProps = {
  className?: string;
  small?: boolean;
};
export const WaitListButton: React.FC<WaitListButtonProps> = ({
  className,
  small,
}) => {
  const ButtonText = "Join the wait list!";

  const button = small ? (
    <Button className={className} radius="xl" sx={{ height: 30 }}>
      {ButtonText}
    </Button>
  ) : (
    <Button className={className} size="lg">
      {ButtonText}
    </Button>
  );

  return (
    <Link
      href="mailto:figissue@camreronbrill.me?cc=nicolas3104@gmail.com&subject=Fig%20Issue%20Early%20Access%20Request&body=Hi Cameron,%0A%0A {replace with desired use-case/any questions} %0A%0A I would like early access to Fig Issue!%0A%0A Thanks!"
      passHref
    >
      {button}
    </Link>
  );
};
