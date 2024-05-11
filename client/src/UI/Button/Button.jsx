import React from "react";
import "./Button.style.css";

export const ButtonType = Object.freeze({
  Fill: "fill",
  Outline: "outline",
  Gradient: "gradient",
  GradientNoGlow: "gradient-no-glow",
});

function Button(props) {
  const style = {
    padding: `${props.pv} ${props.ph}`,
    "whiteSpace": "nowrap",
  };

  if (props.width) {
    style.width = props.width;
  }

  switch (props.type) {
    case ButtonType.Fill:
      return (
        <button className="button-fill" style={style}>
          {props.text}
        </button>
      );
    case ButtonType.Outline:
      return (
        <button className="button-outline" style={style}>
          {props.text}
        </button>
      );
    case ButtonType.Gradient:
      return (
        <button className="button-gradient" style={style}>
          {props.text}
        </button>
      );
    case ButtonType.GradientNoGlow:
      return (
        <button className="button-gradient-no-glow" style={style}>
          {props.text}
        </button>
      );
    default:
      return null;
  }
}

Button.defaultProps = {
  pv: "8px",
  ph: "24px",
};

export default Button;
