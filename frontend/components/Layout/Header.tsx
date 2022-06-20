import React from "react";
import {
  createStyles,
  Menu,
  Center,
  Header as ManHeader,
  Container,
  Group,
  Burger,
  Anchor,
} from "@mantine/core";
import { useBooleanToggle } from "@mantine/hooks";
import { ChevronDown } from "tabler-icons-react";
import Link from "next/link";
import { ToggleTheme } from "../Toggle/ToggleTheme";
import { WaitListButton } from "../WaitList/Button";

const HEADER_HEIGHT = "4vh";

const useStyles = createStyles((theme) => ({
  inner: {
    height: HEADER_HEIGHT,
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
  },

  links: {
    [theme.fn.smallerThan("sm")]: {
      display: "none",
    },
  },

  burger: {
    [theme.fn.largerThan("sm")]: {
      display: "none",
    },
  },

  link: {
    display: "block",
    lineHeight: 1,
    padding: "8px 12px",
    borderRadius: theme.radius.sm,
    textDecoration: "none",
    color:
      theme.colorScheme === "dark"
        ? theme.colors.dark[0]
        : theme.colors.gray[7],
    fontSize: theme.fontSizes.sm,
    fontWeight: 500,

    "&:hover": {
      backgroundColor:
        theme.colorScheme === "dark"
          ? theme.colors.dark[6]
          : theme.colors.gray[0],
    },
  },

  linkLabel: {
    marginRight: 5,
  },
}));

type HeaderProps = {
  links: {
    link: string;
    label: string;
    links?: { link: string; label: string }[];
  }[];
};

export const Header: React.FC<HeaderProps> = ({ links }) => {
  const { classes } = useStyles();
  const [opened, toggleOpened] = useBooleanToggle(false);
  const items = links.map((link) => {
    const menuItems = link.links?.map((item) => (
      <Menu.Item key={item.link}>{item.label}</Menu.Item>
    ));

    if (menuItems) {
      return (
        <Menu
          key={link.label}
          trigger="hover"
          delay={0}
          transitionDuration={0}
          placement="end"
          gutter={1}
          control={
            <Link
              href={link.link}
              className={classes.link}
              onClick={(event) => event.preventDefault()}
            >
              <Center>
                <span className={classes.linkLabel}>{link.label}</span>
                <ChevronDown size={12} />
              </Center>
            </Link>
          }
        >
          {menuItems}
        </Menu>
      );
    }

    return (
      <Link
        key={link.label}
        href={link.link}
        onClick={(event) => event.preventDefault()}
      >
        <Anchor className={classes.link}>{link.label}</Anchor>
      </Link>
    );
  });

  return (
    <ManHeader height={HEADER_HEIGHT} sx={{ borderBottom: 0 }}>
      <Container className={classes.inner} fluid>
        <Group>
          <Burger
            opened={opened}
            onClick={() => toggleOpened()}
            className={classes.burger}
            size="sm"
          />
          <h1
            style={{
              minWidth: "fit-content",
              padding: "10px",
            }}
          >
            Fig Issue
          </h1>
          {/*<LogoHere />*/}
        </Group>
        <Group spacing={5} className={classes.links}>
          {items}
        </Group>
        <Group>
          <WaitListButton small />
          <ToggleTheme />
        </Group>
      </Container>
    </ManHeader>
  );
};
