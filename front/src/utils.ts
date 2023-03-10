import { AxiosError } from "axios";
import axios from "axios"

import { toast } from "react-toastify"

type ErrorMessage = {
  field:string;
  message:string;
}

interface ApiErrorResponse{
  errors:Array<ErrorMessage>
}

export const Error = (e:Error | AxiosError)=>{

  if(axios.isAxiosError<ApiErrorResponse>(e)){

    if(!e.response){
      return toast.error("Response not found")
    }

    const { errors } = e.response?.data;

    const message = errors.map((item:ErrorMessage)=>`${item.field}: ${item.message}`)
    return toast.error(message.join(" "))
  }

}
