// ... (imports)

export class ApplicationParameters extends React.Component<ApplicationParametersProps, ApplicationParametersState> {
    // ...
    private async loadRefs(repo: string) {
        try {
            const project = this.props.application ? this.props.application.spec.project : undefined;
            const refs = await services.repositories.refs(repo, project).toPromise();
            // ...
        } catch (err) {
            // ...
        }
    }
    // ...
}