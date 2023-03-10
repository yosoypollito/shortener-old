import { useAppSelector } from "@redux/hooks"
import { selectLink } from "@redux/reducers/links" 
import { LinkState } from "@src/types"

import { toast } from "react-toastify"

import { Link } from "react-router-dom"

import { BiChevronsLeft, BiChevronsRight, BiCopy } from 'react-icons/bi'

import '@components/links/link.success.styles.css'

import { useAppDispatch } from "@redux/hooks"
import { prevStep } from "@redux/reducers/links"


export default function LinkSuccess(){

  const linkState:LinkState = useAppSelector(selectLink)

  const dispatch = useAppDispatch()

  const copy = async()=>{
    await navigator.clipboard.writeText(linkState.link);
    toast.success("Copied")
  }

  return(
    <>
      <h1>Thanks for use this service</h1>
      <div className="link__creation">
        <div className="link__creation__info">
          <span>
            {linkState.link || "Getting link..."}
          </span>
        </div>
        <div className="link__creation__actions">
          <button className="link__creation__actions__button" onClick={()=>dispatch(prevStep())}>
            <BiChevronsLeft size={20}/> Create 
          </button>
          <button onClick={copy} className="link__creation__actions__button">
            Copy <BiCopy size={20}/>
          </button>
          <a href={linkState.link} className="link__creation__actions__button">
            Go <BiChevronsRight size={20}/>
          </a>
        </div>
      </div>
      <h3>We don&apos;t recommend use this link for production environments. <br/> Currently, we are hosted for free and could have problems with stay online</h3>
      <span>If you want to help us to upgrade this services, consider a <Link to="donate">donation</Link></span>
    </>
  )
}
