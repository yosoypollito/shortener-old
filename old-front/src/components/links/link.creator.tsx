import { useAppSelector } from "@redux/hooks"
import { selectLink } from "@redux/reducers/links"
import { LinkState } from "@src/types"

import LinkForm from "@components/links/link.form"
import LinkSuccess from "@components/links/link.success"

import '@components/links/link.styles.css'

export default function LinkCreator(){

  const linkState:LinkState = useAppSelector(selectLink);

  const steps = [<LinkForm/>, <LinkSuccess/>]
  return(
    <div className="link__wrapper items-center content-center">
      {linkState && steps[linkState.step]}
    </div>
  )
}
