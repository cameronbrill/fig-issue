import React from "react";
import { createStyles, Container, Group, Anchor } from "@mantine/core";

const useStyles = createStyles((theme) => ({
  footer: {
    height: "5vh",
    borderTop: `1px solid ${
      theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
    }`,
    marginTop: "-5vh",
    position: "relative",
    clear: "both",
  },

  inner: {
    height: "5vh",
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
    paddingTop: theme.spacing.xl,
    paddingBottom: theme.spacing.xl,

    [theme.fn.smallerThan("xs")]: {
      flexDirection: "column",
    },
  },

  links: {
    [theme.fn.smallerThan("xs")]: {
      marginTop: theme.spacing.md,
    },
  },
}));

interface FooterProps {
  links: { link: string; label: string }[];
}

export const Footer: React.FC<FooterProps> = ({ links }) => {
  const { classes } = useStyles();
  const items = links.map((link) => (
    <Anchor<"a">
      color="dimmed"
      key={link.label}
      href={link.link}
      onClick={(event) => event.preventDefault()}
      size="sm"
    >
      {link.label}
    </Anchor>
  ));

  return (
    <footer className={classes.footer}>
      <Container className={classes.inner}>
        <h1
          style={{
            minWidth: "fit-content",
            padding: "10px",
          }}
        >
          Fig Issue
        </h1>
        <Group className={classes.links}>{items}</Group>
      </Container>
    </footer>
  );
};
