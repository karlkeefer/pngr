import { ChangeEvent } from 'react';
import { InputOnChangeData, TextAreaProps } from 'semantic-ui-react';

type RunFunc<T> = (promise: Promise<any>, onSuccess?: (data: T) => void, onFailure?: Function) => void

type InputChangeHandler = (event: ChangeEvent<HTMLInputElement>, data: InputOnChangeData) => void
type TextAreaChangeHandler = (event: ChangeEvent<HTMLTextAreaElement>, data: TextAreaProps) => void

type Post = {
  id?: number,
  title: string,
  body: string
}

type User = {
  email: string,
  pass: string
}
