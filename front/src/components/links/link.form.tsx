import '@components/links/link.form.styles.css'

import { BiChevronRight } from 'react-icons/bi'

import { create } from "@redux/reducers/links"

import * as Regex from "@regex"
import { joiResolver } from "@hookform/resolvers/joi"
import Joi from "joi"
import { useForm } from 'react-hook-form'

type FormData = {
  scope:string
}

const schema = Joi.object({
  scope:Joi.string()
  .required()
  .pattern(new RegExp(Regex.url))
  .messages({
    'string.empty':'Link is required',
    'string.pattern.base':'Not valid link'
  })
})

export default function LinkForm(){

  const { register, handleSubmit, formState:{
    errors
  } } = useForm<FormData>({
    resolver: joiResolver(schema)
  });

  const submit = (data:FormData)=>{
    create(data.scope)
  }


  return(
    <form className="link__form" onSubmit={handleSubmit(submit)}>
      <h1>Short Link</h1>
      <div className="link__form__group">
        <label htmlFor="link">
          Link
        </label>
        <div className="link__form__group--submitable">
          <input {...register("scope")}/>
          <button type="submit" title="Create link">
            <BiChevronRight size={20}/>
          </button>
        </div>
        {errors?.scope?.message && (<span>{errors.scope.message}</span>)}
      </div>
      <h3>Paste your link and get a short one ✅</h3>
    </form>
  )
}
