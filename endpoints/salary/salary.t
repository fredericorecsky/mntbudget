<h1> Fixed Salary </h1>

<h2> Add new salary </h2>

<form id="new_salary" method="POST" action="/salary/">
    <ul>
        <li>
            Description:
            <input type="text" name="description" value=""/>
        </li>
        <li>
            Bruto by Month:
            <input type="text" name="bruto_month" value=""/>
        </li>
        <li>
            <input type="checkbox" name="dutch_30" value="1"> has dutch 30% ruling
        </li>
    </ul>
</form>

<h2> My fixed income: </h2>

<ul>
    <li>
        <table>
            <tr>
                <th>Description</th>
                <th>Salary Bruto</th>
                <th>30% rule</th>
                <th>Monthly tax</th>
                <th>Estimated Netto</th>
                <th>By year</th>
                <th>Taxable by year</th>
                <th>Year Tax</th>
            </tr>
            {{ range $i, $val := . }}
            <tr>
                <td>{{ $val.Description }}</td>
                <td>{{ $val.Bruto_month }}</td>
                <td>{{ $val.Dutch_30 }}</td>
                <td>{{ $val.Monthly_tax }}</td>
                <td>{{ $val.Estimated_net }}</td>
                <td>{{ $val.Bruto_year  }}</td>
                <td>{{ $val.Yearly_tax }}</td>
                <tr>{{ $val.Yearly_taxable }}</td>
             </tr>
             {{end}}
        </table>
    </li>
</ul>