import { useMutation } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { toast } from 'sonner';
import { LoginForm, type LoginFormProps } from '@/components/login-form';
import { requestLogin } from './services/user';

function App() {
  const mutation = useMutation({
    mutationFn: requestLogin,
    onSuccess: () => {
      const urlParams = new URLSearchParams(window.location.search);
      let redirectUri = urlParams.get('redirect_uri') ?? '';
      redirectUri = decodeURIComponent(redirectUri);
      if (!/^\/(console|h5)(\/.*)?$/i.test(redirectUri)) {
        redirectUri = '/h5';
      }
      window.location.href = redirectUri;
    },
    onError: (err: AxiosError) => {
      if (err.status === 403) {
        toast.warning('用户名密码错误, 请重新输入', { classNames: { icon: 'text-orange-400' } });
      } else {
        toast.error(`请求错误 ${err.code}`, { description: err.message, richColors: true });
      }
    },
  });

  return (
    <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-sm">
        <LoginForm onSubmit={mutation.mutate as LoginFormProps['onSubmit']} submitting={mutation.isPending} />
      </div>
    </div>
  );
}

export default App;
