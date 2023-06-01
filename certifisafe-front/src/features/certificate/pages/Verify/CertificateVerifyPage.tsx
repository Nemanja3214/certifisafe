import Button from 'components/forms/Button/Button'
import InputField from 'components/forms/InputField/InputField'
import Menu from 'components/navigation/Menu/Menu'
import PageTitle from 'components/view/PageTitle/PageTitle'
import CertificateVerifyPageCSS from './CertificateVerifyPage.module.scss'
import Card from 'components/view/Card/Card'
import { useState } from 'react'
import * as yup from 'yup' 
import { Formik, Form, Field, ErrorMessage } from 'formik';

import Upload from "assets/actions/upload.png"
import Valid from "assets/actions/valid.png"
import Unvalid from "assets/actions/unvalid.png"

const CertificateVerifyPage = () => {
    const [isValid, setValid] = useState<boolean | undefined>(undefined);
    const [certificateId, setCertificateId] = useState("");

    const schema = yup.object().shape({
        "certificate id": yup.number()
        .test("int", "Must be integer", val => {
          return val === undefined ? false : val % 1 == 0;
        }).typeError("certificate id must be an integer")
      })

      const onClick = () => {
        (async function () {
          
        })()
      }

    return (
        <div className={`page pageWithCols pageWithMenu`}>
            <Menu />
            <div>
                <PageTitle title="Verify certificate" description="Check if any certificate is valid." />
                <div className={`${CertificateVerifyPageCSS.block} pageWithCols`}>
                    <Formik
                        initialValues={{
                            certificateId: '',
                        }}
                        validationSchema={schema}
                        onSubmit={values => {
                
                        }}
                    >
                        {({ errors, touched, setFieldValue }) => (
                        <Form>
                            
                            <div>
                                <div className={CertificateVerifyPageCSS.section}>
                                    <b>Verify certificate by ID</b>
                                    <br />
                                    <br />
                                    <Field name="certificate id" component={InputField} className={CertificateVerifyPageCSS.textInput}  usage="Enter Certificate ID" value={certificateId} onChange={(e:React.ChangeEvent<HTMLInputElement>) => {
                                            setCertificateId(e.target.value);
                                            setFieldValue("certificate id", e.target.value);
                                        }}/>
                                    <ErrorMessage name="certificate id" />
                                        
                                    <br />
                                    <small>Certificate ID can be found on the bottom of every certificate. </small>
                                </div>
                                <small>or</small>
                                <div className={CertificateVerifyPageCSS.section}>
                                    <b>Verify certificate by its copy</b>
                                    <br />
                                    <br />
                                    <label htmlFor="file-upload" className={CertificateVerifyPageCSS.fileUpload}>
                                        Upload a copy<img src={Upload} />
                                    </label>
                                    <input id="file-upload" type="file" />
                                </div>
                                <Button submit={"submit"} text={'VERIFY'} onClick={onClick} />
                            </div>
                            <div className={CertificateVerifyPageCSS.card}>
                                {(isValid !== undefined) &&
                                    <Card>
                                        {(isValid === true) &&
                                            <div>
                                                <img src={Valid} />
                                                <h3>Certificate is <br /> VALID !</h3>
                                            </div>

                                        }
                                        {(isValid === false) &&
                                            <div>
                                                <img src={Unvalid} />
                                                <h3>Certificate is <br /> UNVALID !</h3>
                                            </div>

                                        }
                                    </Card>
                                }
                            </div>
                        </Form>
                        )}
                  </Formik>
                </div>
            </div>
        </div>
    )
}

export default CertificateVerifyPage