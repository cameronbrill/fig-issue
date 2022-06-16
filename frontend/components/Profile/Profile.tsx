import { useMemo, useState } from "react";
import { User } from "../../types/user";

type ProfileProps = {
  user?: User;
};

const Profile: React.FC<ProfileProps> = ({ user }) => {
  const [userArr, setUserArr] = useState<any[]>();

  useMemo(() => {
    const p: any[] = [];
    Object.entries(user).forEach((key, val) => {
      p.push(<div>{`${key}: ${val}`}</div>);
      setUserArr(p);
    });
  }, [user]);

  if (!user) return <></>;

  return <div>{userArr}</div>;
};

export default Profile;
