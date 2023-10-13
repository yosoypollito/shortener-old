import styles from '@components/donation/explanation.module.css'

import DonationButton from './donation.button'

export default function Explanation(){

  return(
    <div className={styles.contribute}>
      <h1>Hey 👋 , you want to contribute?</h1>
      <div className={styles.donate}>
        <p>Yes?, well then you can click the button below and check in what you can contribute</p>
        <p>No?, no problem... just keep in mind!</p>
        <DonationButton/>
        <p>All contributes means an exchange of benefits, so don&apos;t care about the amount.</p>
        <p>If you contribute you get a benefit inside the service, just keep in mind use the same email</p>
      </div>
    </div>
  )
}
