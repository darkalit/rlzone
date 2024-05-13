import React from "react";
import "./Registration.style.css";
import "../../UI/GreyContainer/GreyContainer.style.css";
import Button, { ButtonType } from "../../UI/Button/Button";
import { Link, redirect } from "react-router-dom";
import { useState } from 'react'
import Modal from "../../UI/Modal/Modal"
import "../../UI/Modal/Modal.style.css"

import LogoMain from "../../Images/LogoMain.svg";
import Background from "../../Images/MainBG.png"


export default function Registration() {

  const [showModal, setShowModal] = useState(false)

  const closeModal = () => {
    setShowModal(false)
    redirect('/')
  }

  const handlerModalButton = () => {
    setShowModal(false)
    redirect('/')
  }

  const handlerRegistration = () => {
    setShowModal(true)
  }

  return (
    <>
      <main className="main">
        <div className="image-container">
          <img src={Background} alt="bg" className="main-image" />
        </div>
        <section className="wrapper">
          <div className="content-box">
            <form className="container-registration">
              <div className="main-logo">
                <img src={LogoMain} alt="logo form" style={{ height: "36px" }} />
              </div>
              <div className="input-block">
                <label className="above-name">Epic Games ID</label>
                <input className="name-input" name="name" type="text" placeholder="examplename" />
              </div>
              <div className="input-block">
                <label className="above-name">Email</label>
                <input className="name-input" name="email" type="email" placeholder="username@gmail.com" />
              </div>
              <div className="input-block">
                <label className="above-name">Password</label>
                <input className="name-input" name="password" type="password" placeholder="Password" />
              </div>
              <div className="input-block">
                <label className="above-name">Confirm password</label>
                <input className="name-input" name="password" type="password" placeholder="Password" />
              </div>
              <div className="btn-glow">
                {/* <Link to="/" style={{ marginTop: '48px' }}> */}
                <button className="button-gradient" style={{ marginTop: '48px' }} type="button" onClick={handlerRegistration}>Sign up</button>
                {/* <Button text="Sign up" type={ButtonType.Gradient} onClick={handlerRegistration} width={"175px"} /> */}
                {/* </Link> */}
              </div>
              <div className="register">
                <p className="account">Already have an account?</p> <Link to="/login" className="sign-up" href="#">Login!</Link>
              </div>

            </form>
          </div>
        </section>
        {
          showModal ?
            <Modal onClose={closeModal}>
              <div className='modal-text-center'>
                <h2 className='modal-accent-text'>Registration</h2>
                <p className='modal-main-text'>
                  Your account has been successfully registered
                </p>
                <Button text="Ok" type={ButtonType.Fill} onClick={handlerModalButton} width={100} />
              </div>
            </Modal>
            :
            ''
        }
      </main>
    </>
  );
}